import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-article-content-dialog',
  standalone: true,
  imports: [MatDialogModule, MatButtonModule, MatIconModule],
  templateUrl: './article-content-dialog.component.html',
  styleUrls: ['./article-content-dialog.component.scss']
})
export class ArticleContentDialogComponent {
  constructor(
    public dialogRef: MatDialogRef<ArticleContentDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { title: string, contentBody: string }
  ) {}

  onNoClick(): void {
    this.dialogRef.close();
  }
}
