import { Component } from '@angular/core';
import { DirectoryService } from 'src/app/_services/directory.service';
import { MatIcon } from "@angular/material/icon";
import { MatButton } from '@angular/material/button';

@Component({
  standalone: true,
  selector: 'app-directories',
  imports: [MatIcon, MatButton],
  templateUrl: './directories.component.html',
  styleUrl: './directories.component.scss'
})
export class DirectoriesComponent {
    constructor(private directoryService: DirectoryService) {}

    ngOnInit() {
      this.directoryService.getDirectories().subscribe((data) => {
        console.log(data);
      });
    }
}
