import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavigationComponent } from './navigation/navigation.component';
import { NavigationItemComponent } from './navigation/navigation-item/navigation-item.component';
import { SearchComponent } from './search/search.component';
import { FeaturedComponent } from './featured/featured.component';
import { HttpClientModule } from '@angular/common/http';
import { PostPageComponent } from './post-page/post-page.component';
import { HomePageComponent } from './home-page/home-page.component';
import { PostListComponent } from './post-list/post-list.component';
import { PostListItemComponent } from './post-list/post-list-item/post-list-item.component';
import { ClipboardModule } from '@angular/cdk/clipboard';
import { SearchPageComponent } from './search-page/search-page.component';

@NgModule({
  declarations: [
    AppComponent,
    NavigationComponent,
    NavigationItemComponent,
    SearchComponent,
    FeaturedComponent,
    PostPageComponent,
    HomePageComponent,
    PostListComponent,
    PostListItemComponent,
    SearchPageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ClipboardModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
