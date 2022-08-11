import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Post } from 'src/models/post';
import { GetAllPostsResponse, UpvoteResponse } from 'src/models/response';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  getAllPosts(page: number = 1): Observable<GetAllPostsResponse> {
    return this.http.get<GetAllPostsResponse>(`/api/posts?page=${page}`);
  }

  getFeaturedPost(): Observable<Post> {
    return this.http.get<Post>('/api/posts/featured');
  }

  getPostById(id: number): Observable<Post> {
    return this.http.get<Post>(`/api/posts/${id}`)
  }

  upvote(id: number) {
    return this.http.put<UpvoteResponse>(`/api/posts`, { id })
  }
}
