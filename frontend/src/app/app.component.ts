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
    this.httpService
      .getToken(username, password)
      .subscribe((data: LoginResponse) => {
        this.token = data.token;
      });
  }
}
