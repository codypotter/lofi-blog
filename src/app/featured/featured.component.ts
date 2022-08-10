import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Post } from 'src/models/post';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-featured',
  templateUrl: './featured.component.html',
  styleUrls: ['./featured.component.scss']
})
export class FeaturedComponent implements OnInit {

  post?: Observable<Post>;
  
  constructor(private postsService: PostsService) { 
    this.post = this.postsService.getFeaturedPost();
  }

  ngOnInit(): void {
    this.postsService.getFeaturedPost()
  }

}
