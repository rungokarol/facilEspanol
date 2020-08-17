import { TestBed, getTestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';

import { HttpService, LoginResponse } from './http.service';

const username = 'ala';
const password = 'makota';
const loginEndpoint = 'http://localhost:8080/user/login';

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
});
