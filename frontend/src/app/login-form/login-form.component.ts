import { Component } from '@angular/core';
import { HttpService, LoginResponse } from './../services/http.service';

class ErrorMessage {
  constructor(public message: string) {}
}

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss'],
})
export class LoginFormComponent {
  hide = true;
  token: string = null;
  error: ErrorMessage;
  username = '';
  password = '';

  constructor(private httpService: HttpService) {}

  loginHandler() {
    this.error = undefined;
    this.httpService.getToken(this.username, this.password).subscribe(
      (data: LoginResponse) => {
        this.token = data.token;
      },
      (err) => {
        console.log(err.error);
        this.error = new ErrorMessage(err.error);
        this.token = null;
      },
    );
  }
}
// TODO
// 1. unsubscribe -> is it neccessary to keep subscription as a member and unsubscribe in ngOnDestroy?
// 2. rxjs operators
