import {
  async,
  ComponentFixture,
  TestBed,
  getTestBed,
  fakeAsync,
  tick,
} from '@angular/core/testing';
import { FormBuilder } from '@angular/forms';

import { RegisterFormComponent } from './register-form.component';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { HttpService } from '../services/http.service';
import { of, throwError } from 'rxjs';

describe('RegisterFormComponent', () => {
  let component: RegisterFormComponent;
  let fixture: ComponentFixture<RegisterFormComponent>;
  let httpServiceMock: jasmine.SpyObj<HttpService>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [RegisterFormComponent],
      providers: [
        FormBuilder,
        {
          provide: HttpService,
          useValue: jasmine.createSpyObj('HttpService', ['registerUser']),
        },
      ],
      schemas: [NO_ERRORS_SCHEMA],
    }).compileComponents();
  }));

  beforeEach(() => {
    const injector = getTestBed();
    fixture = TestBed.createComponent(RegisterFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    httpServiceMock = injector.get(HttpService) as jasmine.SpyObj<HttpService>;
  });

  it('registerUser calls http service', fakeAsync(() => {
    httpServiceMock.registerUser.and.returnValue(of(undefined));

    component.controls.name.setValue('user');
    component.controls.password.setValue('pass');
    component.controls.email.setValue('dummy@email.com');

    component.registerUser();

    tick();
    expect(httpServiceMock.registerUser).toHaveBeenCalledWith({
      username: 'user',
      password: 'pass',
      email: 'dummy@email.com',
    });
  }));

  it('registerUser handles http service error', fakeAsync(() => {
    httpServiceMock.registerUser.and.returnValue(throwError(`test errror`));

    component.controls.name.setValue('user');
    component.controls.password.setValue('pass');
    component.controls.email.setValue('dummy@email.com');

    component.registerUser();

    tick();
    expect(httpServiceMock.registerUser).toHaveBeenCalledWith({
      username: 'user',
      password: 'pass',
      email: 'dummy@email.com',
    });
  }));

  it('name must be at least 3 characters long', () => {
    const nameControl = component.controls.name;

    nameControl.setValue('xx');
    expect(nameControl.valid).toBeFalsy();
    nameControl.setValue('xxx');
    expect(nameControl.valid).toBeTruthy();
  });

  it('email must have proper format', () => {
    const emailControl = component.controls.email;

    emailControl.setValue('x@');
    expect(emailControl.valid).toBeFalsy();
    emailControl.setValue('x@x.com');
    expect(emailControl.valid).toBeTruthy();
  });

  it('password must be at least 3 characters long', () => {
    const passwordControl = component.controls.password;

    passwordControl.setValue('xx');
    expect(passwordControl.valid).toBeFalsy();
    passwordControl.setValue('xxx');
    expect(passwordControl.valid).toBeTruthy();
  });

  it('password must equal repeat passowrd', () => {
    const passwordControl = component.controls.password;
    const repeatPasswordControl = component.controls.repeatPassword;

    passwordControl.setValue('pass');
    repeatPasswordControl.setValue('notequal');
    expect(repeatPasswordControl.valid).toBeFalsy();

    repeatPasswordControl.setValue('pass');
    expect(repeatPasswordControl.valid).toBeTruthy();
  });
});
