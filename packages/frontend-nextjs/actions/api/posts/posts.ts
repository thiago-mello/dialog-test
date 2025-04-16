"use server";

import { API_MY_POSTS_URL, API_POSTS_URL } from "@/constants/api";
import { deleteRequest, getRequest, postRequest, putRequest } from "../base";

export interface Post {
  id?: string;
  content: string;
  is_public: boolean;
}

export interface PostProjection {
  id: string;
  content: string;
  created_at: string;
  updated_at: string;
  user: UserProjection;
  like_count: number;
  user_liked_this_post: boolean;
}

export interface UserProjection {
  id: string;
  name: string;
  bio?: string;
}

export interface PostsResponse {
  posts: PostProjection[];
  nextCursor?: string;
}

/**
 * Creates a new post
 * @param post - The post object containing content and visibility settings
 * @returns Promise containing the created post
 * @throws Error if post creation fails
 */
export async function saveNewPost(post: Post): Promise<Post> {
  const response = await postRequest<Post>(API_POSTS_URL, post);
  if (response.status >= 400) {
    throw new Error("error while creating post");
  }

  return response.body as Post;
}

/**
 * Retrieves a single post by ID
 * @param id - The unique identifier of the post
 * @returns Promise containing the post if found, undefined otherwise
 * @throws Error if fetching fails (except 404)
 */
export async function getPost(id: string): Promise<Post | undefined> {
  const response = await getRequest<Post>(`${API_POSTS_URL}/${id}`);
  if (response.status === 404) {
    return;
  }
  if (response.status >= 400) {
    throw new Error("Error fetching post");
  }

  return response.body as Post;
}

/**
 * Updates an existing post
 * @param id - The unique identifier of the post to update
 * @param post - The updated post data
 * @returns Promise containing the updated post
 * @throws Error if update fails
 */
export async function updatePost(id: string, post: Post): Promise<Post> {
  const response = await putRequest<Post>(`${API_POSTS_URL}/${id}`, post);
  if (response.status >= 400) {
    throw new Error("Error updating post");
  }
  return response.body as Post;
}

/**
 * Retrieves a list of posts with pagination support
 * @param last_seen_id - Optional cursor for pagination, ID of last post from previous page
 * @param currentUserPosts - If true, fetches only posts by current user
 * @returns Promise containing posts array and next cursor
 * @throws Error if fetching fails
 */
export async function listPosts(
  last_seen_id?: string,
  currentUserPosts: boolean = false
): Promise<PostsResponse> {
  const params = {
    last_seen_id,
  };

  const url = currentUserPosts ? API_MY_POSTS_URL : API_POSTS_URL;
  const response = await getRequest<PostProjection[]>(url, params);
  if (response.status >= 400) {
    throw new Error("Error fetching posts");
  }
  if (!response.body) {
    return { posts: [] };
  }

  const length = response.body.length;
  const nextCursor = length > 0 ? response.body[length - 1].id : undefined;

  return {
    posts: response.body,
    nextCursor,
  };
}

/**
 * Deletes a post by its ID
 * @param id - The unique identifier of the post to delete
 * @returns Promise that resolves when deletion is complete
 * @throws Error if deletion fails
 */
export async function deletePost(id: string): Promise<void> {
  const url = `${API_POSTS_URL}/${id}`;
  await deleteRequest(url);
}
