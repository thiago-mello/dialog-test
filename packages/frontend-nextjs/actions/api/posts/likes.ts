"use server";

import { API_POSTS_URL } from "@/constants/api";
import { deleteRequest, postRequest } from "../base";

/**
 * Adds or removes a like from a post
 * @param postId - The unique identifier of the post
 * @param is_like - Boolean flag indicating whether to like (true) or unlike (false) the post
 * @throws {Error} Throws an error if the request fails (status code is not 204)
 * @returns Promise that resolves when the like/unlike operation completes successfully
 */
export async function likePost(
  postId: string,
  is_like: boolean = true
): Promise<void> {
  const url = `${API_POSTS_URL}/${postId}/likes`;

  const response = is_like
    ? await postRequest(url, undefined)
    : await deleteRequest(url, undefined);

  if (response.status !== 204) {
    throw new Error("Failed to like post");
  }
}
