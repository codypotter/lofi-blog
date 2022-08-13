import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { GetPaginatedPostsResponse } from 'src/models/response';
import { PostsService } from '../posts.service';

@Component({
  selector: 'app-search-page',
  templateUrl: './search-page.component.html',
  styleUrls: ['./search-page.component.scss']
})
export class SearchPageComponent implements OnInit {

  searchResponse?: Observable<GetPaginatedPostsResponse>;
  requestErr = null;

  constructor(private postsService: PostsService, private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    this.route.queryParamMap.subscribe((params) => {
      let category = params.get("category");
      let query = params.get("query");
      this.searchResponse = this.postsService.search(query, category).pipe(
        catchError(err => {
          this.requestErr = err;
          return throwError(err.message);
        })
      );
    });
  }

}
