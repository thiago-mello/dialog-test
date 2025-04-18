// app/timeline/loading.tsx
import { Skeleton } from "@/components/ui/skeleton";

export default function Loading() {
  return (
    <div className="container mx-auto max-w-3xl py-8">
      {/* Page title and new post button */}
      <div className="flex flex-row justify-between items-center mb-8">
        <Skeleton className="w-[180px] h-8 animate-pulse" />
        <Skeleton className="w-[160px] h-8 animate-pulse" />
      </div>

      {[1, 2, 3, 4].map((item) => (
        <div
          key={item}
          className="border rounded-lg p-4 shadow-sm space-y-4 animate-pulse mb-4"
        >
          {/* Name and timestamp */}
          <div className="flex flex-col gap-2">
            <Skeleton className="h-4 w-24" /> {/* Name */}
            <Skeleton className="h-3 w-32" /> {/* Timestamp */}
          </div>

          {/* Post content - variable height */}
          <div className="space-y-2">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-[90%]" />
            {item % 2 === 0 && <Skeleton className="h-4 w-[70%]" />}
          </div>

          {/* Likes & buttons */}
          <div className="flex justify-between items-center pt-2">
            <Skeleton className="h-4 w-20" /> {/* Likes */}
            <div className="flex gap-2">
              <Skeleton className="h-6 w-6 rounded" />
              <Skeleton className="h-6 w-6 rounded" />
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
