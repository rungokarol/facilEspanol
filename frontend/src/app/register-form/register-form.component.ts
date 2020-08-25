import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { HttpService } from '../services/http.service';

@Component({
  selector: 'app-register-form',
  templateUrl: './register-form.component.html',
  styleUrls: ['./register-form.component.scss'],
})
export class RegisterFormComponent implements OnInit {
  registerForm = this.fb.group(
    {
      name: ['ala', [Validators.required, Validators.minLength(3)]],
      email: ['ala@a', [Validators.required, Validators.email]],
      password: ['dupa', [Validators.required, Validators.minLength(3)]],
      repeatPassword: ['dupaaa', Validators.required],
    },
    {
      validator: equal('password', 'repeatPassword'),
    }
  );

  constructor(private fb: FormBuilder, private httpServie: HttpService) {}

  registerUser() {
    console.log(`register`);
    this.httpServie
      .registerUser({
        username: this.registerForm.controls.name.value,
        password: this.registerForm.controls.password.value,
      })
      .subscribe({
        next: (data) => console.log(data),
        error: (err) => console.log(err),
      });
  }

  ngOnInit() {}

  hasError(controlName: string): boolean {
    const control = this.registerForm.controls[controlName];
    return control.invalid && control.dirty;
  }
}

function equal(controlName: string, matchingControlName: string) {
  return (formGroup: FormGroup) => {
    const control = formGroup.controls[controlName];
    const matchingControl = formGroup.controls[matchingControlName];

    if (matchingControl.errors && !matchingControl.errors.mustMatch) {
      // return if another validator has already found an error on the matchingControl
      return;
    }

    // set error on matchingControl if validation fails
    if (control.value !== matchingControl.value) {
      matchingControl.setErrors({ mustMatch: true });
    } else {
      matchingControl.setErrors(null);
    }
  };
}
