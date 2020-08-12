import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';

import { AppMaterialModule } from './app-material/app-material.module';

import { AppComponent } from './app.component';
import { LoginFormComponent } from './login-form/login-form.component';

@NgModule({
  declarations: [AppComponent, LoginFormComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AppMaterialModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
