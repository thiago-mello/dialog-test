import { listPosts, PostsResponse } from "@/actions/api/posts/posts";
import { PostsList } from "./components/posts-list";
import { QueryClientWrapper } from "@/providers/query";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import Link from "next/link";

async function getInitialPosts(): Promise<PostsResponse> {
  return await listPosts();
}

export default async function TimelinePage() {
  const { posts, nextCursor } = await getInitialPosts();

  return (
    <div className="container mx-auto max-w-3xl py-8">
      <div className="flex flex-row justify-between items-center mb-8">
        <h1 className="text-2xl font-bold">Linha do Tempo</h1>
        <Button asChild>
          <Link href="posts/new" className="text-white">
            <Plus />
            Nova Postagem
          </Link>
        </Button>
      </div>

      <QueryClientWrapper>
        <PostsList initialPosts={posts} initialNextCursor={nextCursor} />
      </QueryClientWrapper>
    </div>
  );
}
