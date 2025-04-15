"use server";
import { API_AUTH_URL } from "@/constants/api";
import { SessionData, sessionOptions } from "@/lib/session";
import { getIronSession } from "iron-session";
import { cookies } from "next/headers";

interface ErrorResponse {
  message: string;
  type: "INTERNAL_ERROR" | "USER_NOT_FOUND";
}

export async function handleLogin(
  email: string,
  password: string
): Promise<ErrorResponse | undefined> {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions
  );

  // checks if user is logged in
  if (session.userId) {
    return;
  }

  const options: RequestInit = {
    method: "POST",
    body: JSON.stringify({
      email: email,
      password: password,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  };

  let response;
  try {
    response = await fetch(API_AUTH_URL, options);
  } catch (error: unknown) {
    if (error instanceof Error) {
      if (error.cause instanceof AggregateError) {
        return {
          message: "Error when connecting to API",
          type: "INTERNAL_ERROR",
        };
      }
    }

    return {
      message: "Unkown error when querying user data",
      type: "INTERNAL_ERROR",
    };
  }

  if (response?.status === 404) {
    return {
      message: "User not found",
      type: "USER_NOT_FOUND",
    };
  }

  const userData = await response.json();
  session.userId = userData.user.id;
  session.email = userData.user.email;
  session.name = userData.user.name;
  session.accessToken = userData.access_token;

  await session.save();
}
