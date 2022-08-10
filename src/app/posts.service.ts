import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { Post } from 'src/models/post';

interface getAllPostsResponse {
  page: number;
  more: boolean;
  posts: Post[];
}

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(private http: HttpClient) { }

  getAllPosts(page: number = 1): Observable<getAllPostsResponse> {
    return this.http.get<getAllPostsResponse>(`/api/posts?page=${page}`, {
      headers: {
        'Access-Control-Allow-Origin': '*'
      }
    });
  }

  getFeaturedPost(): Observable<Post> {
    return this.http.get<Post>('/api/posts/featured');
  }
}
