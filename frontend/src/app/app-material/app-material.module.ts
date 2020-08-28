import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatIconModule,
  MatFormFieldModule,
  MatInputModule,
  MatButtonModule,
  MatCardModule,
} from '@angular/material/';

@NgModule({
  declarations: [],
  imports: [CommonModule,
    BrowserAnimationsModule,
    ],
  exports: [
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    BrowserAnimationsModule,
  ],
})
export class AppMaterialModule {}
