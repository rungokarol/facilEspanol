import {
  async,
  TestBed,
  getTestBed,
  fakeAsync,
  tick,
} from '@angular/core/testing';

import { LoginFormComponent } from './login-form.component';
import { AppMaterialModule } from '../app-material/app-material.module';
import { HttpService, LoginResponse } from '../services/http.service';
import { of, asyncScheduler, throwError } from 'rxjs';

describe('LoginFormComponent', () => {
  let injector: TestBed;
  let component: LoginFormComponent;
  let httpServiceMock: jasmine.SpyObj<HttpService>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [LoginFormComponent],
      imports: [AppMaterialModule],
      providers: [
        LoginFormComponent,
        {
          provide: HttpService,
          useValue: jasmine.createSpyObj('HttpService', ['getToken']),
        },
      ],
    });
    injector = getTestBed();
    component = injector.get(LoginFormComponent);
    httpServiceMock = injector.get(HttpService) as jasmine.SpyObj<HttpService>;
  }));

  it('loginHandler stores token received from http service', fakeAsync(() => {
    const loginResp: LoginResponse = {
      token: 'dummy_token',
    };
    httpServiceMock.getToken.and.returnValue(of(loginResp, asyncScheduler));

    component.loginHandler('user', 'pass');
    tick();

    expect(httpServiceMock.getToken).toHaveBeenCalledWith('user', 'pass');
    expect(component.token).toEqual('dummy_token');
  }));

  it('loginHandler does nothing when http service throws error', fakeAsync(() => {
    httpServiceMock.getToken.and.returnValue(throwError('Test Error'));

    component.loginHandler('user', 'pass');
    tick();

    expect(httpServiceMock.getToken).toHaveBeenCalledWith('user', 'pass');
    expect(component.token).toBeNull();
  }));
});
