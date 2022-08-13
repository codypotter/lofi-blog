import { Post } from "./post";

export interface GetPaginatedPostsResponse {
    page: number;
    more: boolean;
    posts: Post[];
}

export interface UpvoteResponse {
    upvotes: number;
}