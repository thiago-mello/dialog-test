"use server";

import { API_REGISTER_USER_URL, API_USER_EXISTS_URL } from "@/constants/api";
import { getUrlQueryStringFromParams } from "@/utils/url";

export interface User {
  name: string;
  email: string;
  bio?: string;
  password: string;
  password_confirm: string;
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
