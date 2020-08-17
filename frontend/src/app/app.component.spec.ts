import {
  TestBed,
  async,
  fakeAsync,
  ComponentFixture,
} from '@angular/core/testing';

import { AppComponent } from './app.component';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from './app-routing/app-routing.module';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
import { LoginFormComponent } from './login-form/login-form.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { AppMaterialModule } from './app-material/app-material.module';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AppComponent', () => {
  let router: Router;
  let location: Location;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule.withRoutes(routes),
        AppMaterialModule,
        HttpClientTestingModule,
      ],
      declarations: [AppComponent, PageNotFoundComponent, LoginFormComponent],
      schemas: [NO_ERRORS_SCHEMA],
    });

    router = TestBed.get(Router);
    location = TestBed.get(Location);

    fixture = TestBed.createComponent(AppComponent);
    router.initialNavigation();
  }));

  it('renders router outlet', () => {
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('router-outlet')).not.toBeNull();
  });

  it('navigates to "" redirects to /login', fakeAsync(() => {
    router.navigate(['']).then(() => {
      expect(location.path()).toBe('/login');
    });
  }));

  it('navigates to /login', fakeAsync(() => {
    router.navigate(['login']).then(() => {
      expect(location.path()).toBe('/login');
    });
  }));
});
