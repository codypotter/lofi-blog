import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { Post } from 'src/models/post';
import { PostsService } from '../posts.service';
import { Clipboard } from '@angular/cdk/clipboard';

@Component({
  selector: 'app-post-page',
  templateUrl: './post-page.component.html',
  styleUrls: ['./post-page.component.scss']
})
export class PostPageComponent implements OnInit {

  post?: Observable<Post>;
  href = '';

  constructor(private postsService: PostsService, private route: ActivatedRoute, private clipboard: Clipboard) { }

  ngOnInit(): void {
    this.post = this.postsService.getPostById(this.route.snapshot.params.id);
    this.href = window.location.href;
  }

  onUpvote(upvoteButton: HTMLElement): void {
    this.postsService.upvote(this.route.snapshot.params.id).subscribe({
      next: (upvoteResponse) => {
        upvoteButton.textContent = `👍${upvoteResponse.upvotes}`
      }
    })
  }

  copyToClipboardWithParameter(shareButton: HTMLElement): void {
    const successful = this.clipboard.copy(window.location.href);
    if (successful) {
      shareButton.textContent = 'copied!';
      setTimeout(function() {
        shareButton.textContent = '📋share';
      }, 3000)
    }
  }
}
