import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginFormComponent } from './../login-form/login-form.component';
import { PageNotFoundComponent } from './../page-not-found/page-not-found.component';
import {RegisterFormComponent} from '../register-form/register-form.component';

export const routes: Routes = [
  {
    path: 'register',
    component: RegisterFormComponent,
  },
  {
    path: 'login',
    component: LoginFormComponent,
  },
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: '**', component: PageNotFoundComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  declarations: [],
})
export class AppRoutingModule {}
