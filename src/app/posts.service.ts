import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Post } from 'src/models/post';
import { GetPaginatedPostsResponse, UpvoteResponse } from 'src/models/response';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  getAll(page: number = 1): Observable<GetPaginatedPostsResponse> {
    return this.http.get<GetPaginatedPostsResponse>(`/api/posts?page=${page}`);
  }

  getFeatured(): Observable<Post> {
    return this.http.get<Post>('/api/posts/featured');
  }

  search(query: string | null, category: string | null): Observable<GetPaginatedPostsResponse> {
    let queryString = '?';
    if (query) {
      queryString += `query=${query}&`;
    }
    if (category) {
      queryString += `category=${category}`;
    }
    return this.http.get<GetPaginatedPostsResponse>(`/api/posts/${queryString}`);
  }

  getById(id: number): Observable<Post> {
    return this.http.get<Post>(`/api/posts/${id}`)
  }

  upvote(id: number) {
    return this.http.put<UpvoteResponse>(`/api/posts`, { id })
  }
}
