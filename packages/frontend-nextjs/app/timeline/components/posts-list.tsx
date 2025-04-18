"use client";
import { useEffect, useRef, useState } from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { PostItem } from "./post";
import { Spinner } from "@/components/ui/spinner";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import {
  listPosts,
  PostProjection,
  PostsResponse,
} from "@/actions/api/posts/posts";
import { CircleAlert } from "lucide-react";
import Link from "next/link";

async function fetchPosts(
  cursor?: string,
  currentUserOnly?: boolean
): Promise<PostsResponse> {
  return await listPosts(cursor, currentUserOnly);
}

interface PostsListProps {
  initialPosts: PostProjection[];
  initialNextCursor?: string;
  currentUserOnly?: boolean;
  currentUserId?: string;
}

export function PostsList({
  initialPosts,
  initialNextCursor,
  currentUserOnly,
  currentUserId,
}: PostsListProps) {
  const observerTarget = useRef<HTMLDivElement>(null);
  const [localPosts, setLocalPosts] = useState<PostProjection[]>(initialPosts);

  const {
    data,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    status,
    error,
  } = useInfiniteQuery({
    queryKey: ["posts"],
    queryFn: ({ pageParam }) =>
      fetchPosts(pageParam as string, currentUserOnly),
    initialPageParam: initialNextCursor,
    getNextPageParam: (lastPage) => lastPage.nextCursor,
    initialData: {
      pages: [{ posts: initialPosts, nextCursor: initialNextCursor }],
      pageParams: [undefined],
    },
    staleTime: Infinity,
  });

  // This effect sets up an Intersection Observer to implement infinite scrolling
  // When the observer target becomes 50% visible in the viewport (threshold: 0.5),
  // and there are more posts to fetch (hasNextPage) and we're not currently fetching (isFetchingNextPage),
  // it triggers the next page fetch
  // The observer is cleaned up when the component unmounts
  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting && hasNextPage && !isFetchingNextPage) {
          fetchNextPage();
        }
      },
      { threshold: 0.5 }
    );

    const currentTarget = observerTarget.current;
    if (currentTarget) {
      observer.observe(currentTarget);
    }

    return () => {
      if (currentTarget) {
        observer.unobserve(currentTarget);
      }
    };
  }, [fetchNextPage, hasNextPage, isFetchingNextPage]);

  // Combine all pages of posts into a flat array
  useEffect(() => {
    if (data) {
      const allPosts = data.pages.flatMap((page) => page.posts);
      setLocalPosts(allPosts);
    }
  }, [data]);

  if (status === "error") {
    return (
      <Alert variant="destructive" className="mt-4">
        <AlertDescription>
          Erro ao carregar posts:{" "}
          {error instanceof Error ? error.message : "Erro desconhecido"}
        </AlertDescription>
      </Alert>
    );
  }

  const handlePostDeleted = (postId: string) => {
    setLocalPosts((currentPosts) =>
      currentPosts.filter((post) => post.id !== postId)
    );
  };

  return (
    <div className="space-y-4">
      {localPosts.length === 0 && (
        <Alert>
          <CircleAlert />
          <AlertTitle>Nenhum post encontrado</AlertTitle>
          <AlertDescription>
            Não há posts para mostrar.
            <Link className="ml-1 underline text-blue-500" href="/posts/new">
              Crie seu primero post.
            </Link>
          </AlertDescription>
        </Alert>
      )}

      {localPosts.map((post) => (
        <PostItem
          key={post.id}
          post={post}
          currentUserId={currentUserId}
          onPostDeleted={handlePostDeleted}
        />
      ))}

      <div ref={observerTarget} className="h-4" />

      {isFetchingNextPage && (
        <div className="flex justify-center py-4">
          <Spinner className="h-6 w-6 text-primary" />
        </div>
      )}

      {!hasNextPage && localPosts.length > 0 && (
        <p className="text-center text-gray-500 py-4">
          Não há mais posts para carregar
        </p>
      )}
    </div>
  );
}
