import { listPosts, PostsResponse } from "@/actions/api/posts/posts";
import { PostsList } from "../components/posts-list";
import { QueryClientWrapper } from "@/providers/query";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import Link from "next/link";
import { getIronSession } from "iron-session";
import { cookies } from "next/headers";
import { SessionData, sessionOptions } from "@/lib/session";
import { redirect } from "next/navigation";

async function getInitialPosts(): Promise<PostsResponse> {
  // list posts for current user only
  return await listPosts(undefined, true);
}

/**
 * Retrieves the current user's ID from the session.
 *
 * @async
 * @function getCurrentUserId
 * @returns {Promise<string | undefined>} The user ID if logged in
 * @throws {Redirect} Redirects to login page with expired=true if user is not logged in
 */
async function getCurrentUserId(): Promise<string | undefined> {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions
  );

  // If user is not logged in, redirect to login page
  if (!session.userId) {
    redirect("/?expired=true");
  }

  return session.userId;
}

export default async function TimelinePage() {
  const { posts, nextCursor } = await getInitialPosts();
  const currentUserId = await getCurrentUserId();

  return (
    <div className="container mx-auto max-w-3xl py-8">
      <div className="flex flex-row justify-between items-center mb-8">
        <h1 className="text-2xl font-bold">Meus Posts</h1>
        <Button asChild>
          <Link href="posts/new" className="text-white">
            <Plus />
            Nova Postagem
          </Link>
        </Button>
      </div>

      <QueryClientWrapper>
        <PostsList
          initialPosts={posts}
          initialNextCursor={nextCursor}
          currentUserId={currentUserId}
          currentUserOnly={true}
        />
      </QueryClientWrapper>
    </div>
  );
}
