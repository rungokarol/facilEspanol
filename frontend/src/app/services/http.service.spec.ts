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

  it('should return error when not ok response received', () => {
    const errorBody = 'User not found';

    service.getToken(username, password).subscribe(
      () => fail('should have failed'),
      (error: string) => {
        expect(error).toEqual(errorBody);
      }
    );

    const req = httpMock.expectOne(loginEndpoint);
    expect(req.request.method).toBe('POST');
    req.flush(errorBody, { status: 404, statusText: 'error happened' });
  });

  it('should return error when unknown error received', () => {
    const progressEvent = new ProgressEvent(`error`);
    const errorBody = { error: progressEvent };

    service.getToken(username, password).subscribe(
      () => fail('should have failed'),
      (error: string) => {
        expect(error).toEqual(
          `Something bad happened; please try again later.`
        );
      }
    );

    const req = httpMock.expectOne(loginEndpoint);
    expect(req.request.method).toBe('POST');
    req.flush(errorBody, { status: 0, statusText: 'Unknown error' });
  });

  it('should call register endpoint', () => {
    const responseBody = null;
    const registerRequest = new RegisterRequest(
      'user',
      'pass',
      'dummy@email.com'
    );
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
