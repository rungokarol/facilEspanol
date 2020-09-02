import {
  TestBed,
  async,
  fakeAsync,
  ComponentFixture,
  tick,
} from '@angular/core/testing';

import { AppComponent } from './app.component';
import { RouterTestingModule } from '@angular/router/testing';
import { routes } from './app-routing/app-routing.module';
import { Router } from '@angular/router';
import { Location } from '@angular/common';
import { AppModule } from './app.module';

describe('AppComponent', () => {
  let router: Router;
  let location: Location;
  let fixture: ComponentFixture<AppComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [RouterTestingModule.withRoutes(routes), AppModule],
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

  it('navigates to /register', fakeAsync(() => {
    router.navigate(['register']);
    tick();
    expect(location.path()).toBe('/register');
  }));

  it('navigates to /unexpected renders PageNotFoundComponent', fakeAsync(() => {
    router.navigate(['/unexpected']).then(() => {
      const compiled = fixture.debugElement.nativeElement;
      expect(compiled.querySelector('app-page-not-found')).not.toBeNull();
      tick();
    });
  }));
});

// TODO
// 1. proper way to test routing - lecture needed
