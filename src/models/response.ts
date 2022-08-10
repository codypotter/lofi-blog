import { Post } from "./post";

export interface GetAllPostsResponse {
    page: number;
    more: boolean;
    posts: Post[];
}