import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.scss']
})
export class SearchComponent implements OnInit {
  
  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  onInputClick() {
    this.router.navigate(['/search']);
  }

  onButtonClick(query: string) {
    if (query == "") {
      return;
    }
    this.router.navigate(['/search'], {
      queryParams: {
        query: encodeURIComponent(query)
      }
    })
  }
}
