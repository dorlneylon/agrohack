// import { DateField } from '@mui/x-date-pickers/DateField';
import dayjs, { Dayjs } from 'dayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DateTimeField } from '@mui/x-date-pickers/DateTimeField';
import { DatePicker } from '@mui/x-date-pickers';

export default function Picker() {
  return (
    <div className="picker">
      <LocalizationProvider dateAdapter={AdapterDayjs}>
      <DemoContainer
        components={['DateTimeField', 'DateTimeField', 'DateTimeField']}
      >
          <DatePicker
              // label="Выборная дата"
              format="L HH"
              value={dayjs('2021-01-06')}
          />
      </DemoContainer>
    </LocalizationProvider>
    </div>
  );
}