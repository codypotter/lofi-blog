import { Post } from "./post";

export interface GetAllPostsResponse {
    page: number;
    more: boolean;
    posts: Post[];
}

export interface UpvoteResponse {
    upvotes: number;
}