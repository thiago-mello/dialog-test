import { SessionData, sessionOptions } from "@/lib/session";
import { getUrlQueryStringFromParams } from "@/utils/url";
import { getIronSession, IronSession } from "iron-session";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export interface ApiResponse<T> {
  status: number;
  body?: T;
}

/**
 * Makes a GET request to the specified path with optional query parameters
 * @param path - The URL path to make the request to
 * @param params - Optional query parameters to append to the URL
 * @returns Promise resolving to the response data of type T, or undefined if not found
 */
export async function getRequest<T>(
  path: string,
  params: any = undefined
): Promise<ApiResponse<T>> {
  const session = await getSession();

  const url = new URL(path);
  if (params) {
    url.search = getUrlQueryStringFromParams(params);
  }

  const options = getRequestOptions("GET", undefined, session);
  const response = await fetch(url, options);

  return await manageResponse<T>(response);
}

/**
 * Makes a PUT request to the specified path with a request body
 * @param path - The URL path to make the request to
 * @param body - The request body to send
 * @returns Promise resolving to the response data
 */
export async function putRequest<T>(
  path: string,
  body: any
): Promise<ApiResponse<T>> {
  const session = await getSession();

  const options = getRequestOptions("PUT", body, session);
  const response = await fetch(path, options);

  return await manageResponse(response);
}

/**
 * Generates request options for fetch API calls
 * @param method - The HTTP method to use
 * @param body - Optional request body
 * @param session - The user's session containing auth token
 * @returns RequestInit object with headers and body
 */
function getRequestOptions(
  method: "GET" | "PUT" | "POST" | "PATCH" | "DELETE",
  body: any,
  session: IronSession<SessionData>
): RequestInit {
  return {
    method: method,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.accessToken}`,
    },
    body: body ? JSON.stringify(body) : undefined,
  };
}

/**
 * Handles the API response and performs appropriate actions based on status codes
 * @param response - The Response object from the fetch API call
 * @returns Promise resolving to an ApiResponse object containing status and optional body
 * @throws Redirects to login page if response status is 401 (Unauthorized)
 */
async function manageResponse<T>(response: Response): Promise<ApiResponse<T>> {
  if (response.status === 401) {
    redirect("/?expired=true");
  }

  if (response?.status === 404) {
    return {
      status: 404,
    };
  }

  return {
    status: response.status,
    body: await response.json(),
  };
}

/**
 * Gets the current user session
 * @returns Promise resolving to the user's session
 * @throws Redirects to login page if no active session
 */
async function getSession(): Promise<IronSession<SessionData>> {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions
  );

  // If user is not logged in, redirect to login page
  if (!session.userId) {
    redirect("/?expired=true");
  }

  return session;
}
