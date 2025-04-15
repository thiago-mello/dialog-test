"use server";

import { API_NEW_POST_URL } from "@/constants/api";
import { postRequest } from "../base";

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
