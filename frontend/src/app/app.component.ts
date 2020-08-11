import { Component } from '@angular/core';
import { Injectable } from '@angular/core';

import { HttpService, LoginResponse } from './services/http.service';

@Injectable()
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  hide = true;
  token: string = null;

  constructor(private httpService: HttpService) {}

  loginHandler(username: string, password: string) {
    this.httpService.getToken(username, password).subscribe({
      next: (data: LoginResponse) => {
        this.token = data.token;
        console.log(this.token);
      },
      error: (err) => console.log(err.error),
    });
  }
}

// TODO
// 1. unsubscribe -> is ti neccessary to keep subscription as a member and unsubscribe in ngOnDestroy?
