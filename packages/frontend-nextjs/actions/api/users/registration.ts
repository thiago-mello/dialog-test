"use server";

import {
  API_MY_USER_URL,
  API_REGISTER_USER_URL,
  API_USER_EXISTS_URL,
} from "@/constants/api";
import { getUrlQueryStringFromParams } from "@/utils/url";
import { getRequest, putRequest } from "../base";

export interface User {
  name: string;
  email: string;
  bio?: string;
  password?: string;
  password_confirm?: string;
}

export interface UserPublicData {
  id: string;
  name: string;
  email: string;
  bio?: string;
}

/**
 * Creates a new user in the system
 * @param newUser - User object containing name, email, bio, password and password confirmation
 * @returns Promise that resolves to an error object if creation fails, or undefined on success
 */
export async function createUser(
  newUser: User
): Promise<{ error?: string } | undefined> {
  const options: RequestInit = {
    method: "POST",
    body: JSON.stringify(newUser),
    headers: {
      "Content-Type": "application/json",
    },
  };

  let response = await fetch(API_REGISTER_USER_URL, options);
  if (response.status === 400) {
    const body = await response.json();
    return { error: body.message };
  }

  return;
}

/**
 * Checks if an email address is available for registration
 * @param email - Email address to check
 * @returns Promise that resolves to true if email is available, false if already taken or invalid
 */
export async function isEmailAvaliable(email: string): Promise<Boolean> {
  const url = new URL(API_USER_EXISTS_URL);
  url.search = getUrlQueryStringFromParams({ email });

  const options: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  };

  let response = await fetch(url, options);
  if (response.status === 400) {
    return false;
  }

  const body = await response.json();
  return !body.exists;
}

/**
 * Retrieves the currently authenticated user's public data
 * @returns Promise that resolves to the user's public data (id, name, email, bio) if authenticated,
 *          or undefined if not authenticated or request fails
 */
export async function getMyUser(): Promise<UserPublicData | undefined> {
  const response = await getRequest<UserPublicData>(API_MY_USER_URL);
  if (response.status === 404) {
    return;
  }

  return response.body;
}

/**
 * Updates the currently authenticated user's information
 * @param user - User object containing updated user data (name, email, bio, password)
 * @returns Promise that resolves to an error object if update fails, or undefined on success
 */
export async function updateMyUser(
  user: User
): Promise<{ error: string } | undefined> {
  const response = await putRequest<any>(API_MY_USER_URL, user);
  if (response.status === 404) {
    return response.body?.message
      ? { error: response.body.message }
      : undefined;
  }

  return;
}
