import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

const url = 'http://localhost:8080';
const loginEndpoint = '/user/login';
const registerEndpoint = '/user/register';

export interface LoginResponse {
  token: string;
}

export class RegisterRequest {
  constructor(
    public username: string,
    public password: string,
    public email: string
  ) {}
}

@Injectable({
  providedIn: 'root',
})
export class HttpService {
  constructor(private http: HttpClient) {}

  getToken(username: string, password: string): Observable<LoginResponse> {
    return this.http
      .post<LoginResponse>(url + loginEndpoint, {
        username,
        password,
      })
      .pipe(catchError(this.handleError));
  }

  private handleError(error: HttpErrorResponse): Observable<never> {
    console.log(error);
    if (typeof error.error === 'string') {
      return throwError(error.error);
    } else {
      return throwError('Something bad happened; please try again later.');
    }
  }

  registerUser(data: RegisterRequest): Observable<any> {
    console.log(`HTTP SERVICE REGISTER: ${JSON.stringify(data)}`);
    return this.http.post<any>(url + registerEndpoint, data);
  }
}

// TODO
// 1. post request options -> what to return and return type
// 2. is decorator correct way to 'provide' service to the module?
