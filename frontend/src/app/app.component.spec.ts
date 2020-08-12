import { TestBed, async } from '@angular/core/testing';

import { AppComponent } from './app.component';
import { Component } from '@angular/core';

@Component({
  selector: 'app-login-form',
  template: '',
})
class StubLoginFormComponent {}

describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [AppComponent, StubLoginFormComponent],
    }).compileComponents();
  }));

  it('should render login form component', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('app-login-form')).toBeDefined();
  });
});
