import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { GetAllPostsResponse } from 'src/models/response';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-post-list',
  templateUrl: './post-list.component.html',
  styleUrls: ['./post-list.component.scss']
})
export class PostListComponent implements OnInit {

  postsResponse?: Observable<GetAllPostsResponse>;

  constructor(private postsService: PostsService) {
    this.postsResponse = this.postsService.getAllPosts(1);
  }

  ngOnInit(): void {
  }

}
