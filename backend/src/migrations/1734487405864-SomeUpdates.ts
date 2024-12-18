import { MigrationInterface, QueryRunner } from "typeorm";

export class SomeUpdates1734487405864 implements MigrationInterface {
    name = 'SomeUpdates1734487405864'

    public async up(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" ADD "thirdPhone" character varying`);
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`ALTER TABLE "users" DROP COLUMN "thirdPhone"`);
    }

}
