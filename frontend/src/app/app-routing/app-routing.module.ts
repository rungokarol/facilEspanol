import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginFormComponent } from './../login-form/login-form.component';
import { PageNotFoundComponent } from './../page-not-found/page-not-found.component';

export const routes: Routes = [
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
