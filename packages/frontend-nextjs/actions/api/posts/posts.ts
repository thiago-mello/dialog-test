"use server";

import { API_NEW_POST_URL } from "@/constants/api";
import { getRequest, postRequest, putRequest } from "../base";

export interface Post {
  id?: string;
  content: string;
  is_public: boolean;
}
export async function saveNewPost(post: Post): Promise<Post> {
  const response = await postRequest<Post>(API_NEW_POST_URL, post);
  if (response.status >= 400) {
    throw new Error("error while creating post");
  }

  return response.body as Post;
}

export async function getPost(id: string): Promise<Post | undefined> {
  const response = await getRequest<Post>(`${API_NEW_POST_URL}/${id}`);
  if (response.status === 404) {
    return;
  }
  if (response.status >= 400) {
    throw new Error("Error fetching post");
  }

  return response.body as Post;
}

export async function updatePost(id: string, post: Post): Promise<Post> {
  const response = await putRequest<Post>(`${API_NEW_POST_URL}/${id}`, post);
  if (response.status >= 400) {
    throw new Error("Error updating post");
  }
  return response.body as Post;
}
