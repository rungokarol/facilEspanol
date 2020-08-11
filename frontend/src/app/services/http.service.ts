import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

const url = 'http://localhost:8080';
const loginEndpoint = '/user/login';
const minLength = 3;

export interface LoginResponse {
  token: string;
}

@Injectable({
  providedIn: 'root',
})
export class HttpService {

  constructor(private http: HttpClient) {}

  getToken(username: string, password: string): Observable<LoginResponse> {
    if (username.length < minLength || password.length < minLength) {
      return throwError(`HTTP_SERVICE: username or password too short`);
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
