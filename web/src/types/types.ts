export interface Event {
    id?: number;
    name: string;
    startDate: Date;
    endDate: Date;
    location: string;
    color: string;
    type: string;
}
export interface Location {
    id: number;
    name: string;
    color: string;
}

export interface DownTimeType {
    id: number;
    name: string;
    color: string;
}

export interface Credentials {
    username: string;
    password: string;
}

