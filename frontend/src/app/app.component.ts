import { Component } from '@angular/core';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
// import { Observable, throwError } from 'rxjs';
// import { catchError, retry } from 'rxjs/operators';

@Injectable()
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.sass'],
})
export class AppComponent {
  hide = true;
  minLength = 3;

  constructor(private http: HttpClient) {}

  loginHandler(username: string, password: string) {
    if (username.length < this.minLength || password.length < this.minLength) {
      console.log(`username or password too short`);
    } else {
      console.log(`performing login procedure`);
      this.http
        .post<any>('http://localhost:8080/user/login', {
          username,
          password,
        })
        .subscribe((data) => {
          console.log(data);
        });
    }
  }
}
