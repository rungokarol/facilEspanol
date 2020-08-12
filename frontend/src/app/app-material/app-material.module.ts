import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  MatIconModule,
  MatFormFieldModule,
  MatInputModule,
  MatButtonModule,
  MatCardModule,
} from '@angular/material/';

@NgModule({
  declarations: [],
  imports: [CommonModule],
  exports: [
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
  ],
})
export class AppMaterialModule {}
