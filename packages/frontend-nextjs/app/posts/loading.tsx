import { Skeleton } from "@/components/ui/skeleton";

export default function Loading() {
  return (
    <div className="max-w-3xl mx-auto mt-10 px-4 space-y-7">
      {/* Page Title */}
      <Skeleton className="h-8 w-1/3" />

      {/* Editor Toolbar */}
      <div className="flex gap-2">
        <Skeleton className="h-10 w-36" />
        <Skeleton className="h-10 w-10" />
        <Skeleton className="h-10 w-10" />
        <Skeleton className="h-10 w-10" />
        <Skeleton className="h-10 w-10" />
        <Skeleton className="h-10 w-10" />
      </div>

      {/* Text Area */}
      <Skeleton className="h-60 w-full rounded-md" />

      {/* Checkbox */}
      <Skeleton className="h-5 w-40" />

      {/* Submit Button */}
      <div className="flex justify-end">
        <Skeleton className="h-10 w-24 rounded-md" />
      </div>
    </div>
  );
}
