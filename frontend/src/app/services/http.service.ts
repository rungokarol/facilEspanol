import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

const url = 'http://localhost:8080';
const loginEndpoint = '/user/login';

export interface LoginResponse {
  token: string;
}

@Injectable({
  providedIn: 'root',
})
export class HttpService {
  minLength = 3;

  constructor(private http: HttpClient) {}

  getToken(username: string, password: string): Observable<LoginResponse> {
    if (username.length < this.minLength || password.length < this.minLength) {
      console.log(`username or password too short`);
    } else {
      return this.http.post<LoginResponse>(url + loginEndpoint, {
        username,
        password,
      });
    }
  }
}

// TODO
// 1. post request options -> what to return and return type
