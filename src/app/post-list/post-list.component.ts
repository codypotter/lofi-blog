import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { GetPaginatedPostsResponse } from 'src/models/response';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-post-list',
  templateUrl: './post-list.component.html',
  styleUrls: ['./post-list.component.scss']
})
export class PostListComponent implements OnInit {

  postsResponse?: Observable<GetPaginatedPostsResponse>;

  constructor(private postsService: PostsService) {
  }

  ngOnInit(): void {
    this.postsResponse = this.postsService.getAll(1);
  }

}
