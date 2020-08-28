import { TestBed, getTestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';

import { HttpService, LoginResponse, RegisterRequest } from './http.service';

const username = 'ala';
const password = 'makota';
const loginEndpoint = 'http://localhost:8080/user/login';
const registerEndpoint = 'http://localhost:8080/user/register';

describe('HttpService', () => {
  let injector: TestBed;
  let service: HttpService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [HttpService],
    });
    injector = getTestBed();
    service = injector.get(HttpService);
    httpMock = injector.get(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should return token in response', () => {
    const loginResponseMock: LoginResponse = {
      token: 'dummy_token',
    };

    service.getToken(username, password).subscribe((data) => {
      expect(data.token).toEqual(loginResponseMock.token);
    });

    const req = httpMock.expectOne(loginEndpoint);
    expect(req.request.method).toBe('POST');
    req.flush(loginResponseMock);
  });

  it('throws error when username is shorter than 3 chars', () => {
    service.getToken('X', password).subscribe({
      error: (err) => {
        expect(err).toBeTruthy();
      },
    });
    httpMock.expectNone(loginEndpoint);
  });

  it('throws error when password is shorter than 3 chars', () => {
    service.getToken(username, 'X').subscribe({
      error: (err) => {
        expect(err).toBeTruthy();
      },
    });
    httpMock.expectNone(loginEndpoint);
  });

  it('should call register endpoint', () => {
    const responseBody = null;
    const registerRequest = new RegisterRequest('user', 'pass');
    service.registerUser(registerRequest).subscribe((data) => {
      expect(data).toBeNull();
    });

    const req = httpMock.expectOne(registerEndpoint);
    expect(req.request.method).toBe('POST');
    req.flush(responseBody);
  });
});

// TODO
// 1. test register error (not data)
