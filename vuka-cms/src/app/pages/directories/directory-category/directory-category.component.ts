import { Component, Inject } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { AnimationItem } from 'lottie-web';
import { AnimationOptions, LottieComponent } from 'ngx-lottie';
import { BaseError } from 'src/app/_models/base-error.model';
import { DirectoryCategory } from 'src/app/_models/directory-category.model';
import { DirectoryService } from 'src/app/_services/directory.service';
import { MatIcon } from '@angular/material/icon';
import { MatButton } from '@angular/material/button';
import { DirectoryEntryFormDialogComponent } from '../../../components/directory-entry-form-dialog/directory-entry-form-dialog.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-directory-category',
  standalone: true,
  imports: [LottieComponent, MatIcon, MatIcon, MatButton],
  templateUrl: './directory-category.component.html',
  styleUrl: './directory-category.component.scss',
})
export class DirectoryCategoryComponent {
  entryId: string = '';
  directoryCategory: DirectoryCategory | undefined;
  snackBarConfig = {
    duration: 3000,
    panelClass: 'snack-bar-container',
  };

  options: AnimationOptions = {
    path: '/lottie/no-data.json',
  };

  constructor(
    private directoryService: DirectoryService,
    private route: ActivatedRoute,
    private snackBar: MatSnackBar,
    private dialog: MatDialog,
  ) {}

  ngOnInit() {
    this.route.paramMap.subscribe((data: ParamMap) => {
      if (this.entryId === data.get('id')) return;
      this.entryId = data.get('id') || '';
      this.loadEntries(this.entryId);
    });
  }

  animationCreated(animationItem: AnimationItem): void {
    console.log(animationItem);
  }

  loadEntries(id: string) {
    this.directoryService.getDirectoryEntries(id).subscribe({
      next: (data: DirectoryCategory) => {
        this.directoryCategory = data;
      },
      error: (error: BaseError) => {
        console.error('Error fetching directory entries:', error);
        this.snackBar.open('There was an error fetching directory');
      },
    });
  }

  onAddEntry() {
    const dialogRef = this.dialog.open(DirectoryEntryFormDialogComponent, {
      data: { 
        categoryName: this.directoryCategory?.name,
        categoryId: this.entryId
      }
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        // Refresh the entries list
        this.loadEntries(this.entryId);
      }
    });
  }
}
