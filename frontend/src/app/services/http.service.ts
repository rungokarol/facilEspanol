import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';

const url = 'http://localhost:8080';
const loginEndpoint = '/user/login';
const registerEndpoint = '/user/register';
const minLength = 3;

export interface LoginResponse {
  token: string;
}

export class RegisterRequest {
  constructor(public username: string, public password: string) {}
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

  registerUser(data: RegisterRequest): Observable<any> {
    console.log(`HTTP SERVICE REGISTER: ${data.username} ${data.password}`);
    return this.http.post<any>(url + registerEndpoint, data);
  }
}

// TODO
// 1. post request options -> what to return and return type
// 2. is decorator correct way to 'provide' service to the module?
