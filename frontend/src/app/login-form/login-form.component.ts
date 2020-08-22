import { Component } from '@angular/core';
import { HttpService, LoginResponse } from './../services/http.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss'],
})
export class LoginFormComponent {
  hide = true;
  token?: string;
  error?: string;
  username = '';
  password = '';

  constructor(private httpService: HttpService) {}

  loginHandler() {
    this.error = undefined;
    this.httpService.getToken(this.username, this.password).subscribe(
      (data: LoginResponse) => {
        this.token = data.token;
      },
      (err: string) => {
        this.error = err;
        this.token = undefined;
      }
    );
  }
}
// TODO
// 1. unsubscribe -> is it neccessary to keep subscription as a member and unsubscribe in ngOnDestroy?
// 2. rxjs operators
// 3. receive error as string not as error in component - is it correct??
