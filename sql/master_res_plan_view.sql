DROP VIEW master_res_plan;
CREATE ALGORITHM = UNDEFINED DEFINER = `root` @`%` SQL SECURITY DEFINER VIEW `master_res_plan` AS
select
  `div`.`name` AS `division`,
  `dept`.`name` AS `department`,
  `group`.`name` AS `group`,
  `users`.`cn` AS `name`,
  `users`.`location` AS `location`,
  `users`.`resource_type` AS `resource_type`,
  `prog`.`name` AS `program`,
  `prog`.`type` AS `program_type`,
  date_format(`a`.`allocation_time`, '%Y') AS `year`,
  date_format(`a`.`allocation_time`, '%m') AS `month`,
  `a`.`allocation` AS `allocation`
from
  (
    (
      (
        (
          (
            `assignments` `a`
            left join `programs` `prog` on((`a`.`program_id` = `prog`.`id`))
          )
          left join `users` on((`a`.`user_id` = `users`.`id`))
        )
        left join `departments` `group` on((`a`.`dept_id` = `group`.`id`))
      )
      left join `departments` `dept` on(
        (
          cast(
            substring_index(
              substring_index(`group`.`level`, '.', 3),
              '.',
              -(1)
            ) as unsigned
          ) = `dept`.`id`
        )
      )
    )
    left join `departments` `div` on(
      (
        cast(
          substring_index(
            substring_index(`dept`.`level`, '.', 2),
            '.',
            -(1)
          ) as unsigned
        ) = `div`.`id`
      )
    )
  )
where
  `a`.status = 1
order by
  `div`.`name`,
  `dept`.`name`,
  `group`.`name`,
  `users`.`cn`,
  `prog`.`name`,
  date_format(`a`.`allocation_time`, '%Y'),
  date_format(`a`.`allocation_time`, '%m')