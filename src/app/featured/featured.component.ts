import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-featured',
  templateUrl: './featured.component.html',
  styleUrls: ['./featured.component.scss']
})
export class FeaturedComponent implements OnInit {

  @Input()
  post = ""
  
  constructor() { }

  ngOnInit(): void {
  }

}
