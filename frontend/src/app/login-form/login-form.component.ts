import { Component } from '@angular/core';
import { HttpService, LoginResponse } from './../services/http.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent {
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
// 1. unsubscribe -> is it neccessary to keep subscription as a member and unsubscribe in ngOnDestroy?

