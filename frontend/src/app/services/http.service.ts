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
  constructor(private http: HttpClient) {}

  getToken(username: string, password: string): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(url + loginEndpoint, {
      username,
      password,
    });
  }
}

// TODO
// 1. post request options -> what to return and return type
// 2. is decorator correct way to 'provide' service to the module?
