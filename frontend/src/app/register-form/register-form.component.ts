import { Component } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { HttpService } from '../services/http.service';

@Component({
  selector: 'app-register-form',
  templateUrl: './register-form.component.html',
  styleUrls: ['./register-form.component.scss'],
})
export class RegisterFormComponent {
  registerForm = this.fb.group(
    {
      name: ['', [Validators.required, Validators.minLength(3)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(3)]],
      repeatPassword: ['', Validators.required],
    },
    {
      validator: equal('password', 'repeatPassword'),
    }
  );

  constructor(private fb: FormBuilder, private httpServie: HttpService) {}

  registerUser() {
    this.httpServie
      .registerUser({
        username: this.registerForm.controls.name.value,
        password: this.registerForm.controls.password.value,
        email: this.registerForm.controls.email.value,
      })
      .subscribe({
        next: (data) => console.log(data),
        error: (err) => console.log(err),
      });
  }

  get controls() {
    return this.registerForm.controls;
  }
}

function equal(controlName: string, matchingControlName: string) {
  return (formGroup: FormGroup) => {
    const control = formGroup.controls[controlName];
    const matchingControl = formGroup.controls[matchingControlName];

    if (matchingControl.errors && !matchingControl.errors.equal) {
      // return if another validator has already found an error on the matchingControl
      return;
    }

    // set error on matchingControl if validation fails
    if (control.value !== matchingControl.value) {
      matchingControl.setErrors({ equal: true });
    } else {
      matchingControl.setErrors(null);
    }
  };
}

