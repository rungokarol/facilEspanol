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

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('registerUser calls http service', fakeAsync(() => {
    httpServiceMock.registerUser.and.returnValue(of(undefined));

    component.registerForm.controls.name.setValue('user');
    component.registerForm.controls.password.setValue('pass');

    component.registerUser();

    tick();
    expect(httpServiceMock.registerUser).toHaveBeenCalledWith({
      username: 'user',
      password: 'pass',
    });
  }));

  it('registerUser handles http service error', fakeAsync(() => {
    httpServiceMock.registerUser.and.returnValue(throwError(`test errror`));

    component.registerForm.controls.name.setValue('user');
    component.registerForm.controls.password.setValue('pass');

    component.registerUser();

    tick();
    expect(httpServiceMock.registerUser).toHaveBeenCalledWith({
      username: 'user',
      password: 'pass',
    });
  }));
});

