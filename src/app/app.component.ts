import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'lofi-blog';
  testPost = '<h1>test post title</h1><p>test post body</p>'
}
